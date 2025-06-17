import Color from "colorjs.io";

function color(color: string, format: "hex" | "oklch" | "rgb"): string {
  try {
    switch (format) {
      case "hex":
        return new Color(color).to("srgb").toString({ format: "hex" });

      case "rgb":
        return new Color(color).to("srgb").toString();

      case "oklch":
        return new Color(color).to("oklch").toString();
    }
  } catch {
    return "ERROR: Invalid or unsupported input.";
  }
}

function main() {
  const formatSelect = document.querySelector("select")!;
  const colorInput = document.querySelector("input")!;
  const output = document.querySelector("output")!;

  const spaces = ["hex", "oklch", "rgb"] as const;
  for (const space of spaces) {
    formatSelect.insertAdjacentHTML(
      "beforeend",
      `<option${space === "oklch" ? " selected" : ""} value="${space}">${space}</option>`,
    );
  }

  const outputColor = color(
    colorInput.value,
    formatSelect.value as "hex" | "oklch" | "rgb",
  );
  output.innerText = outputColor;

  formatSelect.addEventListener("input", () => {
    const outputColor = color(
      colorInput.value,
      formatSelect.value as "hex" | "oklch" | "rgb",
    );
    output.innerText = outputColor;
  });
  colorInput.addEventListener("input", () => {
    const outputColor = color(
      colorInput.value,
      formatSelect.value as "hex" | "oklch" | "rgb",
    );
    output.innerText = outputColor;
  });
}

main();
