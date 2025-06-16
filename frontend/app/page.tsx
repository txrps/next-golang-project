"use client";

import { useEffect } from "react";

export default function Home() {
  useEffect(() => {
    fetch("http://localhost:8080/api/users")
      .then((res) => res.json())
      .then((data) => {
        console.log(data);
      });
  }, []);
  return <>test</>;
}
