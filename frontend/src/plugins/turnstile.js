export const TurnstileKey = import.meta.env.VITE_TURNSTILE_KEY;
export const TurnstileEnabled = TurnstileKey !== undefined;
export const TurnstileReady = new Promise((resolve) => {
  window.onloadTurnstileCallback = resolve;
});

export const loadTurnstile = () => {
  if (window.turnstile) {
    return;
  }

  const script = document.createElement("script");
  script.src =
    "https://challenges.cloudflare.com/turnstile/v0/api.js?onload=onloadTurnstileCallback&render=explicit";
  script.async = true;
  script.defer = true;
  document.head.appendChild(script);
};
