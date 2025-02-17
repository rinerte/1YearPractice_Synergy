import { useState } from 'react'
import './App.css'
import Selector from './components/selector'

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
      <h1 className='pt-[200px] pb-[40px]'>Tourism APP</h1>
      <Selector/>
    </>
  )
}

export default App
