const interpretedLanguages = ["python"];
const allowedFiles = ["application/zip", "application/x-tar","application/gzip"];
const allowedExtensions = [".zip", ".tar", ".tar.gz"];

const formId = "mainForm";
const buildButtonId = "build";
const runButtonId = "run";
const outputId = "output";
const formErrorId = "formError";
const runCmdId = "runCmd";
const compileCmdId = "compileCmd";
const fileId = "file";
const cmdArgsId =  "cmdArgs";

const getLangHandle = "/getSupportedLanguage";
const uploadHandle = "/upload";
const buildHandle = "/build";
const runHandle = "/run";