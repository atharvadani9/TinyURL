import { ThemeProvider } from "@emotion/react";
import CssBaseline from "@mui/material/CssBaseline";
import TinyURL from "./components";
import darkTheme from "./utils/theme";

function App() {
  return (
    <ThemeProvider theme={darkTheme}>
      <CssBaseline />
      <TinyURL />
    </ThemeProvider>
  );
}

export default App;
