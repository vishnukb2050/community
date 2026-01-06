
import 'dart:io';
import 'package:archive/archive.dart';

void main(List<String> args) {
  if (args.length < 2) {
    print('Usage: unzipper <zipfile> <destdir>');
    exit(1);
  }
  final zipFile = args[0];
  final destDir = args[1];
  
  print('Reading $zipFile...');
  final bytes = File(zipFile).readAsBytesSync();
  
  print('Decoding ZIP...');
  final archive = ZipDecoder().decodeBytes(bytes);
  
  print('Extracting to $destDir...');
  for (final file in archive) {
    final path = '$destDir/${file.name}';
    if (file.isFile) {
      final data = file.content as List<int>;
      File(path).createSync(recursive: true);
      File(path).writeAsBytesSync(data);
      // Set executable permission for bin files
      if (path.contains('/bin/') || path.endsWith('sdkmanager') || path.endsWith('adb')) {
         Process.runSync('chmod', ['+x', path]);
      }
    } else {
      Directory(path).createSync(recursive: true);
    }
  }
  print('Done!');
}
