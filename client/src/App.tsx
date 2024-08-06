import { useState } from 'react'
import { Button, Container, Stack } from '@chakra-ui/react'
import Navbar from './components/Navbar'

function App() {
  const [count, setCount] = useState(0)

  return (
    <Stack h='100vh'>
     <Navbar>

    <Container> 
      {/* TodoForm */}
      {/* TodoList */}
    </Container>
     </Navbar>
    <Button onClick={() => setCount( (count) => count + 1)} > Hello  count is {count} </Button>
    </Stack>
  )
}

export default App
