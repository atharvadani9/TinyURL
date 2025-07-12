import { ThemeProvider } from "@emotion/react";
import CloseIcon from "@mui/icons-material/Close";
import CssBaseline from "@mui/material/CssBaseline";
import IconButton from "@mui/material/IconButton";
import { SnackbarProvider, closeSnackbar } from "notistack";
import TinyURL from "./components";
import darkTheme from "./utils/theme";

function App() {
  return (
    <ThemeProvider theme={darkTheme}>
      <CssBaseline />
      <SnackbarProvider
        maxSnack={3}
        anchorOrigin={{
          vertical: "top",
          horizontal: "right",
        }}
        action={(snackbarId) => (
          <IconButton
            size="small"
            aria-label="close"
            color="inherit"
            onClick={() => closeSnackbar(snackbarId)}
          >
            <CloseIcon fontSize="small" />
          </IconButton>
        )}
      >
        <TinyURL />
      </SnackbarProvider>
    </ThemeProvider>
  );
}

export default App;
