import { Box, Button, Flex, Text, useColorMode } from "@chakra-ui/react";
import { FaGithub } from "react-icons/fa";
import { MdHelp } from "react-icons/md";

export default function Footer() {
  const { colorMode } = useColorMode();
  const textColor = colorMode === "light" ? "gray.800" : "gray.200";
  const buttonVariant = colorMode === "light" ? "outline" : "solid";

  return (
    <Box  color={textColor} py={6} px={4}>
      <Box
        bg={colorMode === "light" ? "white" : "gray.900"}
        py={6} px={8}
        rounded="lg"
        shadow="xl"
        textAlign="center"
      >
        <Text fontSize="xl" fontWeight="semibold" mb={2} color={textColor}>
          Find the source code on GitHub
        </Text>
        <Text fontSize="sm" color={colorMode === "light" ? "gray.600" : "gray.400"} mb={6}>
          Contributions are welcome!
        </Text>
        <Flex justifyContent="center" alignItems="center" gap={4}>
          <Button
            as="a"
            href="https://github.com/hadyrashwan/golang-for-node-devs"
            size="sm"
            variant={buttonVariant}
            colorScheme="gray"
            leftIcon={<FaGithub />}
          >
            Check out the repository
          </Button>
          <Button
            as="a"
            href="https://github.com/hadyrashwan/golang-for-node-devs/issues"
            size="sm"
            variant={buttonVariant}
            colorScheme="gray"
            leftIcon={<MdHelp />}
          >
            Open an issue
          </Button>
        </Flex>
      </Box>
      <Text fontSize="sm" textAlign="center" mt={4} color={textColor}>
        Made using React and Chakra UI.
      </Text>
    </Box>
  );
}
