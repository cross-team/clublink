export function validateUsernameFormat(username?: string): string | null {
  const CUSTOM_ALIAS_MAX_LENGTH = 50;

  if (!username) {
    return null;
  }

  if (username.length >= CUSTOM_ALIAS_MAX_LENGTH) {
    return `Custom alias has to contain less than ${CUSTOM_ALIAS_MAX_LENGTH} characters`;
  }

  if (username.indexOf('#') >= 0) {
    return `Custom alias cannot contain the URL fragment character ('#')`;
  }

  return null;
}
