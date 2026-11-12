import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { BrowserRouter, Route, Routes } from "react-router";
import { Provider } from "@/components/ui/provider";
import Login from "./pages/Login.tsx";
import Register from "./pages/Register.tsx";
import Home from "./pages/Home.tsx";
import { Toaster } from "./components/ui/toaster.tsx";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <Provider defaultTheme="light">
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
        </Routes>
      </BrowserRouter>
      <Toaster />
    </Provider>
  </StrictMode>,
);
