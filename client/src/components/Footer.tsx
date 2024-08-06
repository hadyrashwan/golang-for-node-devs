import { Box, Text } from "@chakra-ui/react";

export default function Footer() {
  return (
    <Box color="gray.200" py={6}>
      <Text fontSize="sm" textAlign={["center"]} mb={[2, 0]}>
        Made using React and Chakra UI.
      </Text>
    </Box>
  );
}
