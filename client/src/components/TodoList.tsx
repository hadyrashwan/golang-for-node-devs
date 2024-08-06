import { Flex, Spinner, Stack, Text } from "@chakra-ui/react";
import TodoItem from "./TodoItem";
import { useQuery } from "@tanstack/react-query";

export type Todo = {
    id: number;
    body: string
    completed: boolean
}

const TodoList = () => {

   const {data:todos,isLoading} = useQuery<Todo[]>({
        queryKey: ["todos"],
        queryFn: async () => {
            try{
                const response = await fetch("http://localhost:4500/api/todos");
                const data = await response.json();
                if(!response.ok) throw new Error(  data.error || "Failed to fetch todos");
                return data.todos || [];
            }catch(error){
               console.log(error); 
            }
        },
    })
	return (
		<>
			<Text fontSize={"4xl"} textTransform={"uppercase"} fontWeight={"bold"} textAlign={"center"} my={2} bgGradient='linear(to-l, #04add9, #4FF7D2)'   bgClip='text'>
				Today's Tasks
			</Text>
			{isLoading && (
				<Flex justifyContent={"center"} my={4}>
					<Spinner size={"xl"} />
				</Flex>
			)}
			{!isLoading && todos?.length === 0 && (
				<Stack alignItems={"center"} gap='3'>
					<Text fontSize={"xl"} textAlign={"center"} color={"gray.500"}>
						All tasks completed! 🤞
					</Text>
				</Stack>
			)}
			<Stack gap={3}>
				{todos?.map((todo) => (
					<TodoItem key={todo._id} todo={todo} />
				))}
			</Stack>
		</>
	);
};
export default TodoList;