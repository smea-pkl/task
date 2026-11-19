import { Box, Button, Flex, HStack } from "@chakra-ui/react";
import { Link } from "react-router";

function NavigationBar() {
  const isLogin = Boolean(
    localStorage.getItem("pkl_token") || sessionStorage.getItem("pkl_token"),
  );

  return (
    <Box>
      <Flex mx="auto" mt="4" align="center" justify="flex-end">
        <HStack>
          <Link to="/">
            <Button size="sm">Home</Button>
          </Link>

          <Button size="sm" variant="ghost">
            Catur
          </Button>
          <Button size="sm" variant="ghost">
            Judol
          </Button>

          {isLogin ? (
            <Link to="/profile">
              <Button size="sm" variant="ghost">
                Profile
              </Button>
            </Link>
          ) : (
            <>
              <Link to="/login">
                <Button size="sm" variant="ghost">
                  Login
                </Button>
              </Link>
              <Link to="/register">
                <Button size="sm" variant="ghost">
                  Register
                </Button>
              </Link>
            </>
          )}
        </HStack>
      </Flex>
    </Box>
  );
}

export default NavigationBar;
