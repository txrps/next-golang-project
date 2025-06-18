"use client";

import { httpClient } from "@/clients/http-client";
import { useEffect } from "react";

export default function Home() {
  const GetApi = async () => {
    const user = await httpClient.get("/test");
    console.log("user", user);
  };

  useEffect(() => {
    GetApi();
  }, []);
  return <>test</>;
}
