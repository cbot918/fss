import java.io.*;
import java.net.*;

public class SimpleClient {
    public static void main(String[] args) {
        try (Socket socket = new Socket("localhost", 12345);
             PrintWriter out = new PrintWriter(socket.getOutputStream(), true);
             BufferedReader in = new BufferedReader(new InputStreamReader(socket.getInputStream()))) {
            
            // Send a message to the server
            out.println("Hello, Server!");
            
            // Read response from the server
            String response = in.readLine();
            System.out.println("Received from server: " + response);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
