import { Container, Stack } from '@chakra-ui/react'
import Navbar from './components/Navbar'
import TodoForm from './components/TodoForm'
import TodoList from './components/TodoList'
import Disclaimer from './components/Disclaimer'
import Footer from './components/Footer'

export const BASE_URL = 'https://go-react-todo.netlify.app/.netlify/functions/api'

function App() {

  return (
    <Stack h='100vh'>
     <Navbar/>

    <Container> 
      <TodoForm/>
      <TodoList/>
      <Disclaimer/>
      <Footer/>
    </Container>
    </Stack>
  )
}

export default App
