import type { Interceptor } from "@connectrpc/connect";

/**
 * Interceptor untuk Connect client.
 * Tempat ditaruh log/hook sebelum & sesudah setiap RPC.
 */

/**
 * Logging interceptor:
 * mencatat setiap permintaan RPC (method + url) beserta waktu respons.
 */
export const loggingInterceptor: Interceptor = (next) => async (req) => {
  console.log("[pkl] RPC", req.method, req.url);
  const res = await next(req);
  console.log("[pkl] done", req.method, req.url);
  return res;
};

/**
 * Auth interceptor:
 * otomatis menyertakan Bearer token ke header Authorization bila ada di localStorage.
 */
export const authInterceptor: Interceptor = (next) => async (req) => {
  const token = localStorage.getItem("pkl_token");
  if (token) {
    req.header.set("Authorization", `Bearer ${token}`);
  }
  return next(req);
};
