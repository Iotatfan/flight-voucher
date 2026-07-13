import { HomePage } from './pages/home'

const apiBaseUrl = import.meta.env.VITE_API_BASE_URL ?? ''

function App() {
  return <HomePage apiBaseUrl={apiBaseUrl} />
}

export default App
