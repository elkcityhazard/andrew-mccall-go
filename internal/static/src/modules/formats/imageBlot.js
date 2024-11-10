import Quill from "quill";
const BlockEmbed = Quill.import("blots/block/embed");

class ImageBlot extends BlockEmbed {
  static blotName = "image";
  static tagName = "img";

  static create(value) {
    let node = super.create();
    node.setAttribute("alt", value.alt);
    node.setAttribute("src", value.url || "/" + value.src);
    node.setAttribute("width", value.width);
    node.setAttribute("height", value.height);
    node.setAttribute("loading", "lazy");
    node.setAttribute("decoding", "async");
    node.setAttribute("aria-label", value.alt);
    return node;
  }

  static value(node) {
    return {
      alt: node.getAttribute("alt"),
      url: node.getAttribute("src") || node.getAttribute("url"),
      width: node.getAttribute("width"),
      height: node.getAttribute("height"),
      loading: node.getAttribute("loading", "lazy"),
      decoding: node.getAttribute("decoding", "async"),
      "aria-label": node.getAttribute("aria-label", "alt"),
    };
  }
}

Quill.register(ImageBlot);
