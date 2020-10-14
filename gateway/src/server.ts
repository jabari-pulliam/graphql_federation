import {ApolloGateway, ServiceEndpointDefinition} from "@apollo/gateway";
import {ApolloServer} from "apollo-server-express";
import express from 'express';
import {createServer} from 'http';
import config from 'config';
import compression from 'compression';
import logger from "./logging";

const services = config.get<ServiceEndpointDefinition[]>('Gateway.services');
const gatewayDebug = config.get<boolean>('Gateway.debug');
const serverPort = config.get<number>('Server.port');

logger.info(`NODE_ENV: ${process.env.NODE_ENV}`);

const gateway = new ApolloGateway({
    debug: gatewayDebug,
    serviceList: services
});

async function doHealthCheck() {
    try {
        await gateway.serviceHealthCheck();
    } catch (e) {
        logger.error("Service health check failed, retrying");
        await doHealthCheck();
    }
}

doHealthCheck().then(() => {
    startServer();
});

function startServer() {
    const app = express();
    const server = new ApolloServer({
        gateway,
        playground: true,
        subscriptions: false,
        context: ({req}) => {
            // Copy the headers from the incoming request into the context. This is needed for the gateway to
            //  populate the headers for the downstream request.
            const headers = req.headers;
            return {
                headers
            };
        }
    });

    app.use('/gateway/ping', function (req, res, next) {
        res.send('ok');
    });

    app.use(compression());
    server.applyMiddleware({app, path: '/gateway'});

    const httpServer = createServer(app);
    httpServer.listen({port: serverPort},
        (): void => logger.info(`🚀 GraphQL is now running on port ${serverPort} at /gateway`));
}

