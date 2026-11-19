import { createConnectTransport } from "@connectrpc/connect-web";
import { createClient } from "@connectrpc/connect";
import { UserService } from "@/gen/pkl/users/v1/user_service_pb";
import { authInterceptor, loggingInterceptor } from "./interceptors";

/**
 * Klien user-service yang sudah punya interceptor (logging + auth).
 * Gunakan saja `userClient` dari sini di halaman.
 */
const transport = createConnectTransport({
  baseUrl: "http://localhost:8080",
  interceptors: [loggingInterceptor, authInterceptor],
});

export const userClient = createClient(UserService, transport);
