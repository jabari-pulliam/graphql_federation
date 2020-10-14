import bunyan, {LoggerOptions} from 'bunyan';
import config from 'config';

const options = config.get<LoggerOptions>('Logging');

const logger = bunyan.createLogger({
    name: options.name,
    serializers: {
        req: require('bunyan-express-serializer'),
        res: bunyan.stdSerializers.res,
        err: bunyan.stdSerializers.err
    },
    level: options.level
});

export default logger;
