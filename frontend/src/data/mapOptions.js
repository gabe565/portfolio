export default {
  center: { lat: 35.52, lng: -97.49 },
  options: {
    zoom: 10.8,
    disableDefaultUI: true,
    scrollwheel: false,
    draggable: false,
    disableDoubleClickZoom: true,
    backgroundColor: "#222",
    styles: [
      {
        stylers: [
          {
            hue: "#607d8b",
          },
          {
            saturation: 10,
          },
          {
            invert_lightness: true,
          },
          {
            lightness: 20,
          },
          {
            gamma: 0.7,
          },
        ],
      },
      {
        featureType: "administrative.land_parcel",
        elementType: "labels",
        stylers: [
          {
            visibility: "off",
          },
        ],
      },
      {
        featureType: "poi",
        elementType: "labels.text",
        stylers: [
          {
            visibility: "off",
          },
        ],
      },
      {
        featureType: "poi.business",
        stylers: [
          {
            visibility: "off",
          },
        ],
      },
      {
        featureType: "road",
        elementType: "labels.icon",
        stylers: [
          {
            visibility: "off",
          },
        ],
      },
      {
        featureType: "road.arterial",
        elementType: "labels",
        stylers: [
          {
            visibility: "off",
          },
        ],
      },
      {
        featureType: "road.highway",
        elementType: "labels",
        stylers: [
          {
            visibility: "off",
          },
        ],
      },
      {
        featureType: "road.local",
        stylers: [
          {
            visibility: "off",
          },
        ],
      },
      {
        featureType: "road.local",
        elementType: "labels",
        stylers: [
          {
            visibility: "off",
          },
        ],
      },
      {
        featureType: "transit",
        stylers: [
          {
            visibility: "off",
          },
        ],
      },
    ],
  },
};
