export type PopmartIP = {
  slug: string;
  name: string;
  description: string;
  keyword: string;
  imageUrl: string;
};

export const POPMART_IPS: PopmartIP[] = [
  {
    slug: 'labubu',
    name: 'Labubu',
    description: 'THE MONSTERS universe including Macaron, Have a Seat, and collabs.',
    keyword: 'pop mart labubu',
    imageUrl: 'https://imagedelivery.net/j3cQhHk0Y2fW12f0X0X0Xw/placeholder-5/public',
  },
  {
    slug: 'skullpanda',
    name: 'Skullpanda',
    description: 'Art-driven character line with many themed blind box series.',
    keyword: 'pop mart skullpanda',
    imageUrl: 'https://imagedelivery.net/j3cQhHk0Y2fW12f0X0X0Xw/placeholder-4/public',
  },
  {
    slug: 'dimoo',
    name: 'Dimoo',
    description: 'Dreamy fantasy character line with high collector demand.',
    keyword: 'pop mart dimoo',
    imageUrl: 'https://imagedelivery.net/j3cQhHk0Y2fW12f0X0X0Xw/placeholder-6/public',
  },
  {
    slug: 'hirono',
    name: 'Hirono',
    description: 'Stylized narrative character series from Lang.',
    keyword: 'pop mart hirono',
    imageUrl: 'https://imagedelivery.net/j3cQhHk0Y2fW12f0X0X0Xw/placeholder-3/public',
  },
  {
    slug: 'molly',
    name: 'Molly',
    description: 'Classic Pop Mart IP with broad crossovers and editions.',
    keyword: 'pop mart molly',
    imageUrl: 'https://imagedelivery.net/j3cQhHk0Y2fW12f0X0X0Xw/placeholder-2/public',
  },
  {
    slug: 'crybaby',
    name: 'Crybaby',
    description: 'Emotive character line with trend-driven releases.',
    keyword: 'pop mart crybaby',
    imageUrl: 'https://imagedelivery.net/j3cQhHk0Y2fW12f0X0X0Xw/placeholder-1/public',
  },
  {
    slug: 'pucky',
    name: 'Pucky',
    description: 'Fairy-tale inspired mini worlds and seasonal sets.',
    keyword: 'pop mart pucky',
    imageUrl: 'https://imagedelivery.net/j3cQhHk0Y2fW12f0X0X0Xw/placeholder-7/public',
  },
  {
    slug: 'hacipupu',
    name: 'Hacipupu',
    description: 'Whimsical character line with playful themed drops.',
    keyword: 'pop mart hacipupu',
    imageUrl: 'https://imagedelivery.net/j3cQhHk0Y2fW12f0X0X0Xw/placeholder-8/public',
  },
];

export const findPopmartIPBySlug = (slug: string) =>
  POPMART_IPS.find((ip) => ip.slug === slug);
