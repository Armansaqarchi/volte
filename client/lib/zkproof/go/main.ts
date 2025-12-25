// main.ts
// Assumes wasm_exec.js is loaded and defines global "Go"

// If you bundle wasm_exec.js, import it for its side effects:
import "./wasm_exec.js";

// Tell TypeScript about the global Go constructor provided by wasm_exec.js
declare global {
    // eslint-disable-next-line no-var
    var Go: {
        new (): {
            argv: string[];
            importObject: WebAssembly.Imports;
            run(instance: WebAssembly.Instance): Promise<void> | void;
        };
    };
}

const WASM_PATH = "/proof.wasm";

// Helper to start the Go WASM runtime and call an exported function
export async function runWasm(fnName: string, argv: string[]) {
    const go = new Go();

    // Some wasm_exec versions don't pre-set argv, so we default to []
    go.argv = [...(go.argv || []), ...argv];

    let result: WebAssembly.WebAssemblyInstantiatedSource;

    if ("instantiateStreaming" in WebAssembly) {
        // Modern fast path
        result = await WebAssembly.instantiateStreaming(
            fetch(WASM_PATH),
            go.importObject
        );
    } else {
        // Fallback
        const resp = await fetch(WASM_PATH);
        const bytes = await resp.arrayBuffer();
        result = (await WebAssembly.instantiate(
            bytes,
            go.importObject
        )) as WebAssembly.WebAssemblyInstantiatedSource;
    }

    // Start the Go runtime in the background (do not await)
    Promise.resolve(go.run(result.instance)).then(
        () => console.log("Go runtime exited."),
        (err) => console.error("Go runtime crashed or exited with error:", err)
    );

    // Wait briefly for Go main() to register global JS functions
    await new Promise<void>((resolve) => setTimeout(resolve, 50));

    const fn = (globalThis as any)[fnName] as (() => unknown) | undefined;

    if (typeof fn !== "function") {
        throw new Error(
            `Global function "${fnName}" is not available. Got: ${String(
                fn
            )}. Check that your Go main() doesn't exit and the function is exported correctly.`
        );
    }

    console.log("Still running the function");
    const out = fn();
    console.log("ended running the function");
    return out;
}
