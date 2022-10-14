function GetCrc32(Instr) {
    if (typeof (window.Crc32Table) != "undefined")
        return;
    window.Crc32Table = new Array(256);
    var i, j;
    var Crc;
    for (i = 0; i < 256; i++) {
        Crc = i;
        for (j = 0; j < 8; j++) {
            if (Crc & 1)
                Crc = ((Crc >> 1) & 0x7FFFFFFF) ^ 0xEDB88320;
            else
                Crc = ((Crc >> 1) & 0x7FFFFFFF);
        }
        Crc32Table[i] = Crc;
    }
    if (typeof Instr != "string")
        Instr = "" + Instr;
    Crc = 0xFFFFFFFF;
    for (i = 0; i < Instr.length; i++)
        Crc = ((Crc >> 8) & 0x00FFFFFF) ^ Crc32Table[(Crc & 0xFF) ^ Instr.charCodeAt(i)];
    Crc ^= 0xFFFFFFFF;
    return Crc;
} 