import { Box, Alert, AlertIcon, AlertTitle, Stack } from "@chakra-ui/react";

export default function Disclaimer() {
  return (
    <Box py={4}>
      <Stack spacing={4}>
        <Alert status="info" variant="left-accent" rounded="md">
          <AlertIcon />
          <AlertTitle fontSize="md">
            This platform displays user-generated content publicly.
          </AlertTitle>
        </Alert>
      </Stack>
    </Box>
  );
}