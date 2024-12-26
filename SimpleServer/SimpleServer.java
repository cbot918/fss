import java.io.*;
import java.net.*;

public class SimpleServer {
    public static void main(String[] args) {
        try (ServerSocket serverSocket = new ServerSocket(12345)) {
            System.out.println("Server is running on port 12345...");
            
            // Wait for a client connection
            try (Socket clientSocket = serverSocket.accept();
                 BufferedReader in = new BufferedReader(new InputStreamReader(clientSocket.getInputStream()));
                 PrintWriter out = new PrintWriter(clientSocket.getOutputStream(), true)) {
                
                System.out.println("Client connected.");
                
                // Read message from the client
                String message = in.readLine();
                System.out.println("Received from client: " + message);
                
                // Send response to the client
                out.println("Hello, Client! You said: " + message);
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}