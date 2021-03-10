const path = require("path");
const fs = require("fs");
const os = require("os");
const crypto = require("crypto");
const download = require("download");
const ProgressBar = require("progress");

const pkg = require("../package.json");

function getArch() {
  const arch = os.arch();
  switch (arch) {
    case "ia32":
    case "x32":
      return "386";
    case "x64":
      return "amd64";
    case "arm":
      // @ts-expect-error ignore error
      const armv = process.config.variables.arm_version;

      if (!armv) return "armv7";

      return `armv${armv}`;
    case "mips":
      return "mips_hardfloat";
    case "mips64":
      return "mips64_hardfloat";
    case "mipsel":
      return "mipsle_hardfloat";
    case "mips64el":
      return "mips64le_hardfloat";
    default:
      return arch;
  }
}

function getPlatform() {
  const platform = os.platform();
  switch (platform) {
    case "win32":
      return "windows";
    default:
      return platform;
  }
}

function getDownloadURL(version) {
  const url = `https://github.com/axetroy/gmock/releases/download/${version}/gmock_${getPlatform()}_${getArch()}.tar.gz`;
  return url;
}

async function install(version) {
  const url = getDownloadURL(version);
  const binDir = path.join(__dirname, "..", "download");
  const suffix = os.platform() === "win32" ? ".exe" : "";
  const binFile = path.join(binDir, "gmock" + suffix);

  const cacheHex = crypto.createHash("md5").update(url).digest("hex");
  const cachePath = path.join(os.tmpdir(), cacheHex);

  try {
    fs.unlinkSync(binFile);
  } catch {
    // ignore error
  }

  if (fs.existsSync(cachePath)) {
    console.log(`Found cache in '${cachePath}'`);
    fs.copyFileSync(cachePath, binFile);
  } else {
    console.log(`Downloading '${url}'`);
    const bar = new ProgressBar("[:bar] :percent :etas", {
      complete: "=",
      incomplete: " ",
      width: 20,
      total: 0,
    });

    await download(url, binDir, {
      extract: true,
    })
      .on("response", (res) => {
        bar.total = res.headers["content-length"];
        res.on("data", (data) => bar.tick(data.length));
        res.on("error", (err) => {
          console.error(err);
        });
      })
      .on("error", (err) => {
        console.error(err);
      });

    console.log(`Cache binary file to '${cachePath}'`);

    fs.copyFileSync(binFile, cachePath);
  }

  fs.chmod(binFile, 0x755, (err) => {
    // ignore error
    if (err) {
      console.log("This error will be ignore.");
    }
  });
}

install("v" + pkg.version).catch((err) => {
  throw err;
});
