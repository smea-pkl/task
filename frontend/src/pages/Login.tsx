import { useState } from "react";
import {
  Button,
  Card,
  Field,
  Input,
  Stack,
  Flex,
  Checkbox,
} from "@chakra-ui/react";
import { PasswordInput } from "@/components/ui/password-input";
import { toaster } from "@/components/ui/toaster";
import { createConnectTransport } from "@connectrpc/connect-web";
import { ConnectError, createClient } from "@connectrpc/connect";
import { UserService } from "@/gen/pkl/users/v1/user_service_pb";

const transport = createConnectTransport({
  baseUrl: "http://localhost:8080",
});
const userServiceClient = createClient(UserService, transport);

export default function Login() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [visible, setVisible] = useState(false);
  const [rememberMe, setRememberMe] = useState(false);
  const [isLoading, setIsLoading] = useState(false);

  const login = () => {
    // 1. Tolak jika username atau password kosong
    if (!username.trim() || !password.trim()) {
      toaster.create({
        title: "Login Gagal",
        description: "Username dan Password tidak boleh kosong!",
        type: "error",
      });
      return;
    }

    setIsLoading(true);
    userServiceClient
      .login({ username: username.trim(), password })
      .then((res) => {
        // 2. Pastikan token valid (bukan kosong atau "0")
        if (!res.token || res.token === "0") {
          toaster.create({
            title: "Login Gagal",
            description: "Username atau password salah",
            type: "error",
          });
          return;
        }

        // Simpan token sesuai pilihan Remember me
        if (rememberMe) {
          localStorage.setItem("token", res.token);
        } else {
          sessionStorage.setItem("token", res.token);
        }

        toaster.create({
          title: "Login",
          description: "Login Berhasil",
          type: "success",
        });
      })
      .catch((err) => {
        const connectErr = ConnectError.from(err);
        toaster.create({
          title: "Login Gagal",
          description: connectErr.rawMessage || "Gagal terhubung ke server",
          type: "error",
        });
      })
      .finally(() => {
        setIsLoading(false);
      });
  };

  const handleCancel = () => {
    setUsername("");
    setPassword("");
    setRememberMe(false);
  };

  return (
    <Flex minH="100dvh" align="center" justify="center" p="4">
      <Card.Root maxW="sm" w="full">
        <Card.Header>
          <Card.Title>Login</Card.Title>
          <Card.Description>Isi Username dan Password Anda</Card.Description>
        </Card.Header>
        <Card.Body>
          <Stack gap="4" w="full">
            <Field.Root>
              <Input
                placeholder="Username"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
              />
            </Field.Root>
            <Field.Root>
              <PasswordInput
                placeholder="Password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                visible={visible}
                onVisibleChange={setVisible}
              />
            </Field.Root>
          </Stack>
          <Checkbox.Root
            mt="3"
            checked={rememberMe}
            onCheckedChange={(e) => setRememberMe(!!e.checked)}
          >
            <Checkbox.HiddenInput />
            <Checkbox.Control />
            <Checkbox.Label>Remember me</Checkbox.Label>
          </Checkbox.Root>
        </Card.Body>
        <Card.Footer justifyContent="flex-end">
          <Button variant="outline" onClick={handleCancel}>
            Cancel
          </Button>
          <Button variant="solid" loading={isLoading} onClick={login}>
            Sign in
          </Button>
        </Card.Footer>
      </Card.Root>
    </Flex>
  );
}
