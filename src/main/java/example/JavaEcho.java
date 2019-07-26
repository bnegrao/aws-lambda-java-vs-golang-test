package example;

import com.amazonaws.services.lambda.runtime.LambdaLogger;
import com.amazonaws.services.lambda.runtime.Context;

public class JavaEcho {

    public String echo(String phrase, Context context) {
        LambdaLogger logger = context.getLogger();
        logger.log("received : " + phrase);
        return phrase;
    }
}
