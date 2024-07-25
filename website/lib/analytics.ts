import { AnalyticsBrowser } from "@segment/analytics-next";

const getOptions = () => {
  return {
    context: {
      app: {
        name: "Kardinal",
        version: process.env.NEXT_PUBLIC_LANDING_PAGE_VERSION!,
        build: process.env.NEXT_PUBLIC_CF_PAGES_COMMIT_SHA!,
        // Normally this info is collected automatically by Segment, but
        // it misreports the page path when using Next.js's client side routing
        page: {
          path: window.location.pathname,
          referer: document.referrer,
          search: window.location.search,
          title: document.title,
          url: window.location.href,
        },
      },
    },
  };
};

const segment = AnalyticsBrowser.load(
  {
    writeKey: process.env.NEXT_PUBLIC_SEGMENT_WRITE_KEY!,
  },
  {
    disableClientPersistence: true, // we dont like cookies around these parts
  },
);

// Server-side rendered placeholder
let url = new URL("https://example.com");
if (typeof window !== "undefined") {
  // Client-side-only code
  url = new URL(window.location.href);
}
const params = new URLSearchParams(url.search);
type UTMParameter =
  | "utm_source"
  | "utm_medium"
  | "utm_campaign"
  | "utm_term"
  | "utm_content";

const utmParameters: UTMParameter[] = [
  "utm_source",
  "utm_medium",
  "utm_campaign",
  "utm_term",
  "utm_content",
];

for (const param of utmParameters) {
  // If the parameter exists in the URL, store it in localStorage
  if (params.has(param)) {
    localStorage.setItem(param, params.get(param)!);
  }
}

type EventName =
  | "MODAL_OPENED"
  | "MODAL_CLOSED"
  | "BUTTON_CLICK"
  | "NAVIGATE_TO_DOCS"
  | "FORM_SUBMIT"
  | "SCROLLED_1VH"
  | "TIME_ON_PAGE";

interface Payload {
  // identify payload
  email?: string;

  // BUTTON_CLICK payload
  analyticsId?: string; // identifier of the element the user clicked

  // FORM_SUBMIT payload
  formType?: string; // which form is the user submitting
  commPreference?: string;
  features?: string[];

  // TIME_ON_PAGE payload
  duration_seconds?: number;
  page_path?: string;
}

const getUtmParams = (): Record<UTMParameter, string> => {
  // Create an object to store the UTM parameters
  const utmData: Record<UTMParameter, string> = {
    utm_source: "",
    utm_medium: "",
    utm_campaign: "",
    utm_term: "",
    utm_content: "",
  };

  // Iterate over the UTM parameters
  for (const param of utmParameters) {
    // If the parameter exists in localStorage, retrieve it and add it to the utmData object
    if (localStorage.getItem(param)) {
      utmData[param] = localStorage.getItem(param)!; // this might be null but thats ok
    }
  }
  return utmData;
};

// export a simplified interface that ensures options context is always included
const analytics = {
  track: (eventName: EventName, payload?: Payload) => {
    const args = [
      eventName,
      {
        ...payload,
        ...getUtmParams(),
      } || {},
      getOptions(),
    ];

    if (process.env.NODE_ENV === "development") {
      console.debug("TRACK:", ...args);
      return;
    }

    // @ts-ignore
    segment.track(...args);
  },
  identify: (payload: Payload) => {
    const args = [
      {
        ...payload,
        ...getUtmParams(),
      },
      getOptions(),
    ];

    if (process.env.NODE_ENV === "development") {
      console.debug("IDENTIFY:", ...args);
      return;
    }

    segment.identify(...args);
  },
  page: () => {
    const args = [getUtmParams(), getOptions()];
    if (process.env.NODE_ENV === "development") {
      console.debug("PAGE:", ...args);
      return;
    }

    segment.page(...args);
  },
  getUtmParams,
};

export default analytics;
