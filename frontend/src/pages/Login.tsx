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

function Login() {
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
              <Input />
            </Field.Root>
            <Field.Root>
              <Field.Label>Password</Field.Label>
              <PasswordInput />
            </Field.Root>
          </Stack>
          <Checkbox.Root mt="2" value="remember me">
            <Checkbox.HiddenInput />
            <Checkbox.Control />
            <Checkbox.Label>Remember me</Checkbox.Label>
          </Checkbox.Root>
        </Card.Body>

        <Card.Footer justifyContent="flex-end">
          <Button variant="outline">Cancel</Button>
          <Button variant="solid">Sign in</Button>
        </Card.Footer>
      </Card.Root>
    </Flex>
  );
}
export default Login;
