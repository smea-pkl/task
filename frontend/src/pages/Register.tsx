import { useState } from "react";
import { Button, Card, Center, Input } from "@chakra-ui/react";
import { PasswordInput } from "@/components/ui/password-input";
import { toaster } from "@/components/ui/toaster";
import { createConnectTransport } from "@connectrpc/connect-web";
import { ConnectError, createClient } from "@connectrpc/connect";
import { UserService } from "@/gen/pkl/users/v1/user_service_pb";

const transport = createConnectTransport({
  baseUrl: "http://localhost:8080",
});
const userServiceClient = createClient(UserService, transport);

export default function RegisterPage() {
  const [name, setName] = useState<string>("");
  const [username, setUsername] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [isLoading, setIsLoading] = useState<boolean>(false);

  const register = () => {
    if (!name.trim() || !username.trim() || !password.trim()) {
      toaster.create({
        title: "Register",
        description: "Tidak boleh kosong",
        type: "error",
      });
      return;
    }
    const payload = {
      nama: name.trim(),
      password,
      username: username.trim(),
    };
    setIsLoading(true);
    userServiceClient
      .register(payload)
      .then(() => {
        toaster.create({
          title: "Register",
          description: "Register Berhasil",
          type: "success",
        });
        setUsername("");
        setName("");
        setPassword("");
      })
      .catch((err) => {
        const connectErr = ConnectError.from(err);
        toaster.create({
          title: "Register",
          description: connectErr.rawMessage || "Register Gagal",
          type: "error",
        });
      })
      .finally(() => {
        setIsLoading(false);
      });
  };

  return (
    <Center h="100dvh">
      <Card.Root maxW="sm">
        <Card.Body gap="2">
          <Card.Title mt="2">Register</Card.Title>
          <Card.Description>Silahkan register.</Card.Description>
          <Input
            placeholder="Name"
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
          <Input
            placeholder="Username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
          <PasswordInput
            placeholder="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </Card.Body>
        <Card.Footer justifyContent="flex-end">
          <Button
            variant="outline"
            size="sm"
            loading={isLoading}
            onClick={register}
          >
            Register
          </Button>
        </Card.Footer>
      </Card.Root>
    </Center>
  );
}
