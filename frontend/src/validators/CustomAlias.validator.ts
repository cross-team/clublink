export function validateCustomAliasFormat(customAlias?: string): string | null {
  const CUSTOM_ALIAS_MAX_LENGTH = 50;
  const CUSTOM_ALIAS_MIN_LENGTH = 3;

  if (!customAlias) {
    return null;
  }

  if (customAlias.length >= CUSTOM_ALIAS_MAX_LENGTH) {
    return `Custom alias has to contain less than ${CUSTOM_ALIAS_MAX_LENGTH} characters`;
  }

  if (customAlias.length < CUSTOM_ALIAS_MIN_LENGTH) {
    return `Custom alias has to contain at least ${CUSTOM_ALIAS_MIN_LENGTH} characters`;
  }

  if (customAlias.indexOf('#') >= 0) {
    return `Custom alias cannot contain the URL fragment character ('#')`;
  }

  return null;
}
