/*
Constants for interpreted languages, allowed file types,
server endpoints and all HTML element Ids.
 */

// Supported interpreted languages.
const interpretedLanguages = ["python"];
// Allowed submission file types.
const allowedFiles = ["application/zip", "application/x-tar","application/gzip"];
// Allowed submission file extensions.
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