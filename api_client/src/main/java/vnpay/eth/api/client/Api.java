package vnpay.eth.api.client;


import java.io.*;
import java.io.File;
import java.io.FileInputStream;
import java.security.KeyStore;

import javax.net.ssl.SSLContext;

import org.apache.http.HttpEntity;
import org.apache.http.client.methods.CloseableHttpResponse;
import org.apache.http.client.methods.HttpGet;
import org.apache.http.client.methods.HttpPost;
import org.apache.http.client.methods.HttpPut;

import org.apache.http.conn.ssl.SSLContexts;
import org.apache.http.conn.ssl.SSLConnectionSocketFactory;
import org.apache.http.conn.ssl.TrustSelfSignedStrategy;
import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClients;
import org.apache.http.util.EntityUtils;
import org.apache.http.conn.ssl.SSLContextBuilder;
import java.security.KeyStore;
import java.io.FileInputStream;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import org.apache.http.client.config.RequestConfig;
import org.apache.http.client.config.CookieSpecs;

import org.apache.http.NameValuePair;
import org.apache.http.message.BasicNameValuePair;
import org.json.JSONObject;

import org.apache.http.client.entity.UrlEncodedFormEntity;
import org.apache.http.client.methods.HttpRequestBase;

import java.util.*;

public class Api {

    Configuration cfg ;
    CloseableHttpClient httpclient;
    RequestConfig rcfg;
    String rootUrl;

    public  Api(Configuration cfg) {
        this.cfg = cfg;
        this.rootUrl = cfg.getConnection().getRootUrl();
    }
    public void Init(){
      System.out.println("Initialize connection");
      if (rcfg == null) {
         rcfg =  RequestConfig.custom().setCookieSpec(CookieSpecs.STANDARD).build();
      }
      try {
          if (cfg.getConnection().getProtocol().equals("https")) {
            System.out.println("Using SSL connection:  " + cfg.getConnection().getProtocol());

            SSLContextBuilder builder = new SSLContextBuilder();

            Path path = Paths.get(cfg.getConnection().getKeyStoreFile());
            if (path != null) {

                //Add self trusted certificate
                System.out.println("Using self-certificate");
                KeyStore trustStore = KeyStore.getInstance(KeyStore.getDefaultType());
                String password = cfg.getConnection().getKeyStorePassword();
                try (InputStream is = Files.newInputStream(path)) {
                  trustStore.load(is, password.toCharArray());
                }
                builder.loadTrustMaterial(trustStore);
            }
            else{
                //Use strusted all server policy
                System.out.println("Using trusted policy");
                builder.loadTrustMaterial(null, new TrustSelfSignedStrategy());
            }

            SSLConnectionSocketFactory sslsf = new SSLConnectionSocketFactory(
                    builder.build());
            httpclient = HttpClients.custom()
                    .setDefaultRequestConfig(rcfg)
                    .setSSLSocketFactory(sslsf).build();
          } else {
            httpclient = HttpClients.custom()
                    .setDefaultRequestConfig(rcfg).build();
          }
      }catch(Exception ex){
            System.out.println("Error in initilizing connection: " + ex.toString());
      }
  }
  public JSONObject Query(String method, String url, Map paras) {
        System.out.println("Query with method: " + method + ", url: " + url);

        HttpRequestBase  request = null;

        switch(method) {
          case "POST":
              //Post request
              HttpPost requestPost = new HttpPost(url);
              //Parse params
              List<NameValuePair> params = new ArrayList<NameValuePair>();

              List<String> list_key = new ArrayList<String>(paras.keySet());
              for (String key: list_key) {
                 String value = (String) paras.get(key);
                 System.out.println("Add key: " + key + ", value: "+  value);
                 params.add(new BasicNameValuePair(key, value));
              }
              try{
                requestPost.setEntity(new UrlEncodedFormEntity(params,"UTF-8"));
              }catch(Exception ex){
                System.out.println("Cannot Post Parameter");
                return null;
              }

              request = (HttpRequestBase) requestPost;
              break;
          case "GET":
              //Get request
              request = new HttpGet(url);
              break;
          case "PUT":
              request = new HttpPut(url);
              break;
          default:
              System.out.println("Cannot support method: " + method);
              return null;
        }
        request.setHeader("Accept", "application/json");

        //Try parse result and return
        CloseableHttpResponse response = null;
        try {
            //Query data
            response = httpclient.execute(request);

            System.out.println("Status: "  + response.getStatusLine());
            int statusCode = response.getStatusLine().getStatusCode();
            HttpEntity entity = response.getEntity();
            if (entity != null && statusCode == 200 ){
                  String res = EntityUtils.toString(entity);
                  //System.out.println("Response:" + res);
                  JSONObject obj = new JSONObject(res);
                  //System.out.println("Response:" + obj);
                  return obj;
            }
            else{
               System.out.println("No entity");
            }

            //EntityUtils.consume(entity);
        }catch(Exception ex){
           System.out.println("Error to get data: " + ex.toString());
        }
        finally {
              //response.close();
        }
        return null;
  }
  public JSONObject Login(){
        String url = rootUrl + "/login";

        System.out.println("Start login with url = " + url);
        Map<String, String> params = new HashMap<String, String>();
        params.put("username", cfg.getJwt().getUsername());
        params.put("password", cfg.getJwt().getHashPassword());
        return Query("POST",url,params);
  }

  /************ Account functions ************/
  public JSONObject AccountNew(){
      String url = rootUrl + "/api/v1/account/new";

      System.out.println("New Account with url = " + url);
      Map<String, String> params = new HashMap<String, String>();

      return Query("POST",url,params);
  }
  public JSONObject AccountTotal(){
      String url = rootUrl + "/api/v1/account/total";

      System.out.println("Account Total with url = " + url);
      Map<String, String> params = new HashMap<String, String>();
      return  Query("GET",url,params);
  }
}
