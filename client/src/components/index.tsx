import CopyIcon from "@mui/icons-material/ContentCopy";
import Button from "@mui/material/Button";
import Grid from "@mui/material/Grid";
import IconButton from "@mui/material/IconButton";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";
import { enqueueSnackbar } from "notistack";
import { useState, type JSX } from "react";
import { postAPI } from "../utils/httputils";

const TinyURL = (): JSX.Element => {
  const [originalURL, setOriginalURL] = useState("");
  const [shortURL, setShortURL] = useState("");

  const shortenURL = async () => {
    if (
      !originalURL.startsWith("http://") &&
      !originalURL.startsWith("https://")
    ) {
      enqueueSnackbar("Invalid URL", { variant: "error" });
      return;
    }
    const payload = { originalUrl: originalURL };

    const resp = await postAPI<
      typeof payload,
      { shortUrl?: string; error?: string }
    >("/create.tinyurl", payload);
    if (resp.error && resp.error !== "") {
      enqueueSnackbar(`Error shortening URL: ${resp.error}`, {
        variant: "error",
      });
    } else {
      setShortURL("http://localhost:5173/" + resp.shortUrl || "");
      enqueueSnackbar("URL shortened successfully!", { variant: "success" });
    }
  };

  return (
    <Grid
      container
      justifyContent="center"
      sx={{
        pt: 2,
        alignItems: "center",
      }}
      direction={"column"}
    >
      <Grid size={{ xs: 12 }}>
        <Typography variant="h3" sx={{ textAlign: "center" }}>
          {"TinyURL"}
        </Typography>
      </Grid>
      <Grid container sx={{ pt: 2, alignItems: "center", gap: 2 }}>
        <Grid>
          <TextField
            label="Original URL"
            variant="standard"
            value={originalURL}
            onChange={(e) => setOriginalURL(e.target.value)}
          />
        </Grid>
        <Grid sx={{ pt: 1 }}>
          <Button variant="contained" onClick={shortenURL}>
            {"Shorten"}
          </Button>
        </Grid>
      </Grid>
      {shortURL && (
        <Grid container sx={{ pt: 2, alignItems: "center", gap: 2 }}>
          <Grid>
            <Typography variant="h6" sx={{ alignItems: "center" }}>
              {"The shortened URL is: " + shortURL}
            </Typography>
          </Grid>
          <Grid>
            <IconButton
              onClick={() => {
                navigator.clipboard.writeText(shortURL);
                enqueueSnackbar("Copied to clipboard!", { variant: "success" });
              }}
            >
              <CopyIcon />
            </IconButton>
          </Grid>
        </Grid>
      )}
    </Grid>
  );
};

export default TinyURL;
