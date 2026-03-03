export const sanitizePopmartTitle = (value: string) =>
  value
    .replace(/\bAUTHENTIC\b/gi, '')
    .replace(/Opens in a new window or tab/gi, '')
    .replace(/\s+/g, ' ')
    .trim();

const normalize = (value: string) => sanitizePopmartTitle(value);

type SeriesRule = {
  series: string;
  patterns: RegExp[];
};

const SERIES_RULES: SeriesRule[] = [
  {
    series: 'THE MONSTERS - Exciting Macaron Vinyl Face',
    patterns: [
      /\bLABUBU\b.*\bMACARON\b/i,
      /\bEXCITING\s+MACARON\b/i,
      /\bMACARON\s+V\d+\b/i,
    ],
  },
  {
    series: 'THE MONSTERS x One Piece Series',
    patterns: [/\bTHE\s+MONSTERS\b.*\bONE\s*PIECE\b/i, /\bLABUBU\b.*\bONE\s*PIECE\b/i],
  },
  {
    series: 'POP BEAN Pajama Cross Dressing Series',
    patterns: [/\bPAJAMA\s+CROSS\s+DRESSING\b/i],
  },
  {
    series: 'THE MONSTERS x Coca-Cola Series',
    patterns: [/\bTHE\s+MONSTERS\b.*\bCOCA[\s-]?COLA\b/i, /\bLABUBU\b.*\bCOCA[\s-]?COLA\b/i],
  },
  {
    series: 'THE MONSTERS x How to Train Your Dragon Series',
    patterns: [/\bHOW\s+TO\s+TRAIN\s+YOUR\s+DRAGON\b/i],
  },
  {
    series: 'THE MONSTERS x SpongeBob Series',
    patterns: [/\bSPONGE\s*BOB\b/i],
  },
  {
    series: 'THE MONSTERS x Kow Yokoyama Series',
    patterns: [/\bKOW\s+YOKOYAMA\b/i],
  },
  {
    series: 'THE MONSTERS - Have a Seat Series',
    patterns: [/\bHAVE\s+A\s+SEAT\b/i],
  },
  {
    series: 'THE MONSTERS - Big into Energy Series',
    patterns: [/\bBIG\s+INTO\s+ENERGY\b/i],
  },
];

const slugify = (value: string) =>
  sanitizePopmartTitle(value)
    .toLowerCase()
    .replace(/&/g, ' and ')
    .replace(/[^a-z0-9]+/g, '-')
    .replace(/^-+|-+$/g, '');

export const POPMART_SERIES = SERIES_RULES.map((rule) => rule.series);

export const seriesToSlug = (series: string) => slugify(series);

export const seriesFromSlug = (slug: string, candidates: string[]) => {
  const normalizedSlug = slugify(slug);
  if (!normalizedSlug) {
    return '';
  }
  return candidates.find((candidate) => slugify(candidate) === normalizedSlug) ?? '';
};

export const classifyPopmartSeries = (title: string) => {
  const cleaned = normalize(title ?? '');
  if (!cleaned) {
    return '';
  }

  for (const rule of SERIES_RULES) {
    if (rule.patterns.some((pattern) => pattern.test(cleaned))) {
      return rule.series;
    }
  }

  const withoutPrefix = cleaned.replace(/^(authentic\s+)?(pop\s*mart\s+)?/i, '');
  const withoutCharacter = withoutPrefix.replace(
    /^(labubu|the monsters|skullpanda|dimoo|hirono|molly|crybaby|pucky|hacipupu)\s+/i,
    '',
  );
  return withoutCharacter || cleaned;
};
