import { Box, Button, Flex, Text } from "@chakra-ui/react";
import { FaGithub } from "react-icons/fa";
import { MdHelp } from "react-icons/md";

export default function Footer() {
  return (
    <Box color="gray.200" py={4} px={6}>
        <Box
          bg="gray.800"
          py={4} px={6}
          rounded="lg"
          shadow="xl"
          textAlign="center"
        >
          <Text fontSize="xl" fontWeight="semibold" mb={2}>
            Find the source code on GitHub
          </Text>
          <Text fontSize="sm" color="gray.400" mb={6}>
            Contributions are welcome!
          </Text>
          <Flex justifyContent="center" alignItems="center" gap={4}>
            <Button
              as="a"
              href="https://github.com/hadyrashwan/golang-for-node-devs"
              size="sm"
              colorScheme="gray"
              leftIcon={<FaGithub />}
            >
              Check out the repository
            </Button>
            <Button
              as="a"
              href="https://github.com/hadyrashwan/golang-for-node-devs/issues"
              size="sm"
              colorScheme="gray"
              leftIcon={<MdHelp />}
            >
              Open an issue
            </Button>
          </Flex>
        </Box>
      <Text fontSize="sm" textAlign="center" mt={4}>
        Made using React and Chakra UI.
      </Text>
    </Box>
  );
}
