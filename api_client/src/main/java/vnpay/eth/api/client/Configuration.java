package vnpay.eth.api.client;

import static java.lang.String.format;

import java.util.Date;
import java.util.List;
import java.util.Map;

import java.io.IOException;
import java.io.InputStream;
import java.nio.file.Files;
import java.nio.file.Paths;
import org.yaml.snakeyaml.Yaml;

import java.nio.charset.StandardCharsets;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;


public class Configuration {
    public static class  Connection {
        private String host;
        private String port;
        private String protocol;
        private String keyStoreFile;
        private String keyStorePassword;

        public Connection() {
        }

        public String getRootUrl(){
          return this.protocol + "://"+ this.host + ":" + this.port;
        }
        public String getHost() {
          return this.host;
        }

        public void setHost(String host) {
            this.host = host;
        }

        public String getPort() {
          return this.port;
        }

        public void setPort(String port) {
            this.port = port;
        }
        public String getProtocol() {
          return this.protocol;
        }

        public void setProtocol(String protocol) {
            this.protocol = protocol;
        }
        public String getKeyStoreFile() {
          return this.keyStoreFile;
        }

        public void setKeyStoreFile(String keyStoreFile) {
            this.keyStoreFile = keyStoreFile;
        }
        public String getKeyStorePassword() {
          return this.keyStorePassword;
        }

        public void setKeyStorePassword(String keyStorePassword) {
            this.keyStorePassword = keyStorePassword;
        }

        @Override
        public String toString() {
            return new StringBuilder()
                .append( format( "Host: %s\n", host ) )
                .append( format( "Port: %s\n", port ) )
                .append( format( "Protocol: %s\n", protocol ) )
                .append( format( "keyStoreFile: %s\n",keyStoreFile ) )
                .append( format( "keyStorePassword: %s\n",keyStorePassword ) )
                .toString();
        }
    }
    public static class  Jwt {
        private String signkey;
        private String username;
        private String password;

        public Jwt() {
        }

        @Override
        public String toString() {
            StringBuilder sb =  new StringBuilder()
                .append( format( "signkey: %s\n", signkey ) )
                .append( format( "username: %s\n", username ) )
                .append( format( "password: %s\n", password ) );
            return sb.toString();
        }

        public String getHashPassword(){
          try{
            MessageDigest md = MessageDigest.getInstance("MD5");
            byte[] hashInBytes = md.digest(this.password.getBytes(StandardCharsets.UTF_8));

            StringBuilder sb = new StringBuilder();
            for (byte b : hashInBytes) {
               sb.append(String.format("%02x", b));
            }

            return sb.toString();
          }catch(Exception ex){
            System.out.println("Error in hash password: " + ex.toString());
            return "";
          }

        }
        public String getSignkey() {
          return this.signkey;
        }

        public void setSignkey(String signkey) {
            this.signkey = signkey;
        }
        public String getUsername() {
          return this.username;
        }

        public void setUsername(String username) {
            this.username = username;
        }
        public String getPassword() {
          return this.password;
        }

        public void setPassword(String password) {
            this.password = password;
        }

    }
    private String version;
    private Date released;
    private Connection connection;
    private Jwt jwt;

    private static Configuration _instance;

    public static Configuration loadConfig(String fileName){
          if (_instance != null ){
              return Configuration._instance ;
          }
          else{
            Yaml yaml = new Yaml();
            try( InputStream in = Files.newInputStream( Paths.get(fileName) ) ) {
               _instance =  yaml.loadAs(in,Configuration.class);
            }
            catch(Exception ex){
                  System.out.println("Error: " + ex.toString());
                  _instance  = null;
            }
          }
          return Configuration._instance;
    }
    public Configuration() {
    }


    public String getVersion() {
      return this.version;
    }

    public void setVersion(String version) {
        this.version = version;
    }
    public Date getReleased() {
      return this.released;
    }

    public void setReleased(Date released) {
        this.released = released;
    }
    public Connection getConnection() {
      return this.connection;
    }

    public void setConnection(Connection connection) {
        this.connection = connection;
    }
    public Jwt getJwt() {
      return this.jwt;
    }

    public void setJwt(Jwt jwt) {
        this.jwt = jwt;
    }
    @Override
    public String toString() {
        return new StringBuilder()
            .append( format( "Version: %s\n", version ) )
            .append( format( "Released: %s\n", released ) )
            .append( format( "Connection: %s\n", connection.toString() ) )
            .append( format( "Jwt: %s\n", jwt.toString()  ) )
            .toString();
    }
}
