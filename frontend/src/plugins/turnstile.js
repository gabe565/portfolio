import { noop, useScriptTag } from "@vueuse/core";

export const TurnstileKey = import.meta.env.VITE_TURNSTILE_KEY || "1x00000000000000000000AA";
export const TurnstileEnabled = TurnstileKey !== undefined;
export const TurnstileReady = new Promise((resolve) => {
  window.onloadTurnstileCallback = resolve;
});

export const useTurnstile = (onLoaded = noop, options = {}) =>
  useScriptTag(
    "https://challenges.cloudflare.com/turnstile/v0/api.js?onload=onloadTurnstileCallback&render=explicit",
    onLoaded,
    options,
  );
