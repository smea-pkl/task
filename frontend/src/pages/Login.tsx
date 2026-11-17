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
import { createClient } from "@connectrpc/connect";
import { UserService } from "@/gen/pkl/users/v1/user_service_pb";

const transport = createConnectTransport({
  baseUrl: "http://localhost:8080",
});
const userServiceClient = createClient(UserService, transport);

export default function Login() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [visible, setVisible] = useState(false);

  const login = () => {
    userServiceClient
      .login({ username, password })
      .then(() => {
        toaster.create({
          title: "Login",
          description: "Login Berhasil",
          type: "success",
        });
      })
      .catch(() => {
        toaster.create({
          title: "Login",
          description: "Login Gagal",
          type: "error",
        });
      });
  };

  return (
    <Flex minH="100dvh" align="center" justify="center" p="4">
      <Card.Root maxW="sm">
        <Card.Header>
          <Card.Title>Login</Card.Title>
          <Card.Description>Isi Username dan Password Anda</Card.Description>
        </Card.Header>
        <Card.Body>
          <Stack gap="4" w="full">
            <Field.Root>
              <Field.Label>Username</Field.Label>
              <Input
                value={username}
                onChange={(e) => setUsername(e.target.value)}
              />
            </Field.Root>
            <Field.Root>
              <Field.Label>Password</Field.Label>
              <PasswordInput
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                visible={visible}
                onVisibleChange={setVisible}
              />
            </Field.Root>
          </Stack>
          <Checkbox.Root mt="2" value="remember me">
            <Checkbox.HiddenInput />
            <Checkbox.Control />
            <Checkbox.Label>Remember me</Checkbox.Label>
          </Checkbox.Root>
        </Card.Body>
        <Card.Footer justifyContent="flex-end">
          <Button variant="outline" onClick={() => {}}>
            Cancel
          </Button>
          <Button variant="solid" onClick={login}>
            Sign in
          </Button>
        </Card.Footer>
      </Card.Root>
    </Flex>
  );
}
