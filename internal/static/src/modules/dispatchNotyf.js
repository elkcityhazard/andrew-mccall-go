import { Notyf } from "notyf";
import "../../node_modules/notyf/notyf.min.css";

const notyf = new Notyf();

const getNotyf = function () {
  return notyf;
};

export { getNotyf };
