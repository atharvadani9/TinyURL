import { type JSX, useCallback, useEffect } from "react";
import { useParams } from "react-router";
import { postAPI } from "../utils/httputils";

const Redirect = (): JSX.Element => {
  const { param } = useParams();

  const getURL = useCallback(async () => {
    const payload = { shortUrl: param };

    const resp = await postAPI<
      typeof payload,
      { originalUrl?: string; error?: string }
    >("/get.tinyurl", payload);
    if (resp.error && resp.error !== "") {
      console.log(`Error getting URL: ${resp.error}`);
      window.location.href = "http://localhost:5173/404";
    } else {
      window.location.href = resp.originalUrl || "";
    }
  }, [param]);

  useEffect(() => {
    getURL();
  }, [getURL]);

  return <div>Redirecting...</div>;
};

export default Redirect;
