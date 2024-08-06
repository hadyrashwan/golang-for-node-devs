import { Box, Alert, AlertIcon, AlertTitle } from "@chakra-ui/react";

export default function Disclaimer() {
  return (
    <Box py={4} px={6}>
      <Alert status="info" variant="left-accent" rounded="md">
        <AlertIcon />
        <AlertTitle fontSize="md">
          This platform displays user-generated content publicly. <br />
          The platform owner is not responsible for the content's accuracy,
          legality, or appropriateness.
        </AlertTitle>
      </Alert>
    </Box>
  );
}
