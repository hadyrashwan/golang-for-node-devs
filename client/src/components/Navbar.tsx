import { Box, Flex, Button, useColorModeValue, useColorMode, Text, Container, VisuallyHidden } from "@chakra-ui/react";
import { IoMoon } from "react-icons/io5";
import { LuSun } from "react-icons/lu";

export default function Navbar() {
	const { colorMode, toggleColorMode } = useColorMode();

	return (
		<Container maxW={"900px"}>
			<Box bg={useColorModeValue("gray.400", "gray.700")} px={4} my={4} borderRadius={"5"}>
				<Flex h={16} alignItems={"center"} justifyContent={"space-between"}>
					<Flex
						justifyContent={"center"}
						alignItems={"center"}
						gap={3}
						display={{ base: "none", sm: "flex" }}
					>
						<img src='/golang.png' alt='logo' width={50} height={50} />
						<Text fontSize={"40"}>+</Text>
						<img src='/turso.png' alt='logo' width={40} height={40} />
					</Flex>

					<Flex alignItems={"center"} gap={3}>
						<Text fontSize={"lg"} fontWeight={500}>
							Daily Tasks
						</Text>
						<Button onClick={toggleColorMode} variant="ghost" role="switch" aria-checked={colorMode === "light"} title="Toggle Dark Mode">
							{colorMode === "light" ? (
								<>
									<VisuallyHidden>Dark Mode</VisuallyHidden>
									<IoMoon />
								</>
							) : (
								<>
									<VisuallyHidden>Light Mode</VisuallyHidden>
									<LuSun size={20} />
								</>
							)}
						</Button>
					</Flex>
				</Flex>
			</Box>
		</Container>
	);
}