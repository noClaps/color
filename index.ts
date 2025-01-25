import Color from "colorjs.io";

const spaces = Object.keys(Color.spaces);
spaces.push("rgb", "hex");
spaces.sort();
const formats = spaces.join("\n").trim();

const usage = `USAGE: color <color> <format> [--list-formats]`;
const help = `
USAGE: color <color> <format> [--list-formats]

ARGUMENTS:
  <color>               The color that you would like to convert.
  <format>              The format that you would like to convert to.

OPTIONS:
  --list-formats, -f    List all the available formats and exit.
  --help, -h            Display this help message and exit.
`.trim();

function color(args: string[]): string {
  const initialColor = args.slice(0, -1).join(" ");
  const outputFormat = args[args.length - 1];

  try {
    switch (outputFormat) {
      case "hex":
        return new Color(initialColor).to("srgb").toString({ format: "hex" });

      case "rgb":
        return new Color(initialColor).to("srgb").toString();

      default:
        return new Color(initialColor).to(outputFormat).toString();
    }
  } catch {
    return `${Bun.color("red", "ansi")}ERROR:${Bun.color("white", "ansi")} Invalid or unsupported input. Check "color --list-formats" to see if this format is supported.`;
  }
}

const args = process.argv.slice(2);

if (args.includes("-h") || args.includes("--help")) {
  console.log(help);
  process.exit(0);
}

if (args.includes("--list-formats") || args.includes("-f")) {
  console.log(formats);
  process.exit(0);
}

for (const arg of args) {
  if (arg.startsWith("-")) {
    console.log(
      `${Bun.color("red", "ansi")}ERROR:${Bun.color("white", "ansi")} ${arg} is not a valid option. Check "color --help" to see available options.`,
    );
    process.exit(1);
  }
}

if (args.length < 2) {
  console.log(usage);
  process.exit(1);
}

console.log(color(args));
