import {
  Button,
  Card,
  Field,
  Input,
  Stack,
  Flex,
} from "@chakra-ui/react";
import { PasswordInput } from "@/components/ui/password-input";
import { toaster } from "@/components/ui/toaster";
import { UserService } from "@/gen/pkl/users/v1/user_service_pb";
import { useState } from "react";
import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";

const transport = createConnectTransport({
  baseUrl: "http://localhost:8080",
});

const userServiceClient = createClient(UserService, transport);

export default function () {

  const [nama, setNama] = useState();
  const [username, setUsername] = useState();
  const [password, setPassword] = useState();



const register = () => {
 const payload = {
            name: nama,
            password: password,
            username: username
        }

        userServiceClient.
            register(payload).
            then(()=>{
                toaster.create({
                    title: "Register",
                    description: "Register Berhasil",
                    type: "success",
                    
                })
            }).
            catch(()=>{
                toaster.create({
                    title: "Register",
                    description: "Register Gagal",
                    type: "error",
                    
                })
            })
          }

function Register() {
  return (
    <Flex minH="100dvh" align="center" justify="center" p="4">
      <Card.Root maxW="sm">
        <Card.Header>
          <Card.Title>Register</Card.Title>
          <Card.Description>Isi Username dan Password Anda</Card.Description>
        </Card.Header>
        <Card.Body>
          <Stack gap="4" w="full">
            <Field.Root>
              <Field.Label>Nama</Field.Label>
              <Input value={nama} onChange={e=>setNama(e.target.value as any)}/>
            </Field.Root>
            <Field.Root>
              <Field.Label>Username</Field.Label>
              <Input value={username} onChange={e=>setUsername(e.target.value as any)}/>
            </Field.Root>
            <Field.Root>
              <Field.Label>Password</Field.Label>
              <PasswordInput value={password} onChange={e=>setPassword(e.target.value as any)}/>
            </Field.Root>
          </Stack>
        </Card.Body>
        <Card.Footer justifyContent="flex-end">
          <Button variant="outline">Cancel</Button>
          <Button variant="solid" onClick={register}>Register</Button>
        </Card.Footer>
      </Card.Root>
    </Flex>
  );
}
export default Register;
