import { useEffect, useState } from "react";
import { Button, Card, Container, Stack, Text } from "@chakra-ui/react";
import { useNavigate } from "react-router";
import Layout from "@/layout/Layout";
import { userClient } from "@/lib/user-client";

export default function Profile() {
  const navigate = useNavigate();
  const [userId, setUserId] = useState<number | null>(null);
  const [nama, setNama] = useState("");
  const [username, setUsername] = useState("");
  const [error, setError] = useState("");

  useEffect(() => {
    const token =
      localStorage.getItem("pkl_token") || sessionStorage.getItem("pkl_token");
    if (!token) {
      navigate("/login");
      return;
    }
    userClient
      .getProfile({})
      .then((res) => {
        setUserId(res.userId);
        setNama(res.nama);
        setUsername(res.username);
      })
      .catch(() => setError("Gagal memuat profil / token tidak valid"));
  }, [navigate]);

  const logout = () => {
    localStorage.removeItem("pkl_token");
    sessionStorage.removeItem("pkl_token");
    navigate("/login");
  };

  return (
    <Layout>
      <Container maxW="md" mt="10">
        <Card.Root>
          <Card.Header>
            <Card.Title>Profile</Card.Title>
          </Card.Header>
          <Card.Body>
            <Stack gap="3">
              <Text>Username: {username || "-"}</Text>
              <Text>Nama: {nama || "-"}</Text>
              <Text>User ID: {userId ?? "-"}</Text>
              {error && <Text color="red.500">{error}</Text>}
            </Stack>
          </Card.Body>
          <Card.Footer justifyContent="flex-end">
            <Button colorPalette="red" onClick={logout}>
              Logout
            </Button>
          </Card.Footer>
        </Card.Root>
      </Container>
    </Layout>
  );
}
