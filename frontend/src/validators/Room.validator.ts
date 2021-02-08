export function validateRoomFormat(room?: string): string | null {
  const CUSTOM_ALIAS_MAX_LENGTH = 255;

  if (!room) {
    return null;
  }

  if (room.length >= CUSTOM_ALIAS_MAX_LENGTH) {
    return `Custom alias has to contain less than ${CUSTOM_ALIAS_MAX_LENGTH} characters`;
  }

  return null;
}
