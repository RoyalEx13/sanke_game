import { serve } from "bun";

serve({
  fetch(req) {
    const url = new URL(req.url);
    
    if (url.pathname === "/") {
      return new Response(Bun.file("./index.html"), {
        headers: { "Content-Type": "text/html" },
      });
    }

    if (url.pathname.endsWith(".html")) {
      return new Response(Bun.file("./index.html"), {
        headers: { "Content-Type": "text/html" },
      });
    }
    if (url.pathname.endsWith(".js")) {
      return new Response(Bun.file("./wasm_exec.js"), {
        headers: { "Content-Type": "application/javascript" },
      });
    } else if (url.pathname.endsWith(".wasm")) {
      return new Response(Bun.file("../snake_game.wasm"), {
        headers: { "Content-Type": "application/wasm" },
      });
    } else {
      return new Response("404: Not Found", { status: 404 });
    }
  },
  port: 3000,
});

