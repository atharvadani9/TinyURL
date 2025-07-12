import CloseIcon from "@mui/icons-material/Close";
import CssBaseline from "@mui/material/CssBaseline";
import IconButton from "@mui/material/IconButton";
import { ThemeProvider } from "@mui/material/styles";
import { SnackbarProvider, closeSnackbar } from "notistack";
import { useRoutes } from "react-router-dom";
import routes from "./router";
import darkTheme from "./utils/theme";

function App() {
  const content = useRoutes(routes);

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
        {content}
      </SnackbarProvider>
    </ThemeProvider>
  );
}

export default App;
