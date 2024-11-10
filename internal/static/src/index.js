import { editorInit } from "./modules/editorConfig.js";
import { initNotificationRemoval } from "./modules/removeNotification.js";

async function init() {
  try {
    editorInit();
    initNotificationRemoval();
  } catch (err) {
    throw new Error(err);
  }
}

document.addEventListener("DOMContentLoaded", init);
