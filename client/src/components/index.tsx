import Button from "@mui/material/Button";
import Grid from "@mui/material/Grid";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";
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
      console.log("Invalid URL");
      return;
    }
    const payload = { originalUrl: originalURL };

    const resp = await postAPI<
      typeof payload,
      { shortUrl?: string; error?: string }
    >("/create.tinyurl", payload);
    if (resp.error && resp.error !== "") {
      console.error("Error shortening URL:", resp.error);
    } else {
      setShortURL(resp.shortUrl || "");
    }
  };

  return (
    <Grid
      container
      justifyContent="center"
      sx={{
        px: 2,
        alignItems: "center",
      }}
      direction={"column"}
    >
      <Grid size={{ xs: 12 }}>
        <Typography variant="h3" sx={{ textAlign: "center" }}>
          {"TinyURL"}
        </Typography>
      </Grid>
      <Grid container sx={{ pt: 2 }}>
        <Grid>
          <TextField
            label="Original URL"
            variant="outlined"
            value={originalURL}
            onChange={(e) => setOriginalURL(e.target.value)}
          />
        </Grid>
        <Grid>
          <Button variant="contained" onClick={shortenURL}>
            {"Shorten"}
          </Button>
        </Grid>
      </Grid>
      <Grid container sx={{ pt: 2 }}>
        <Grid>
          <Typography variant="h6" sx={{ textAlign: "center" }}>
            {"The shortened URL is: " + shortURL}
          </Typography>
        </Grid>
      </Grid>
    </Grid>
  );
};

export default TinyURL;
