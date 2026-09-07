
public class LogLevels {
    
    public static String message(String logLine) {
      String message = logLine.split("]: ")[1];
      return message.trim();
    }

    public static String logLevel(String logLine) {
      int start = logLine.indexOf("[");
      int end = logLine.indexOf("]");
      return logLine.substring(start + 1, end).toLowerCase();
    }

    public static String reformat(String logLine) {
      String message = LogLevels.message(logLine);
      String level = LogLevels.logLevel(logLine);

      return message + " (" + level + ")";
    }
}
