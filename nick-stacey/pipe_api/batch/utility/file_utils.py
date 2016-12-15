from __future__ import unicode_literals
import hashlib
import os
import errno
import codecs
import zipfile
import shutil
import glob
import subprocess

from chardet.universaldetector import UniversalDetector
from ftfy import fix_file


def get_hash(file_handler):
    md5 = hashlib.md5()
    for data in iter(lambda: file_handler.read(8192), b''):
        md5.update(data)
    return md5.hexdigest()


def detect_encoding(in_path, encoding='utf8'):
    file_handler = open(in_path, 'r')
    detector = UniversalDetector()
    sample = file_handler.readlines(300)
    for line in sample:
        detector.feed(line)
        if detector.done:
            break
    detector.close()
    result = detector.result
    if result['encoding']:
        encoding = result['encoding']
    return encoding


def fix_encoding(in_path, in_encoding, out_encoding='utf8'):
    """Attempt to clean up some of the more common encoding screw-ups."""
    in_fh = codecs.open(in_path, 'r+', in_encoding, errors='ignore')
    in_name = in_fh.name
    tmp_name = os.path.join(os.path.dirname(in_fh.name), 'converting.tmp')
    out_fh = codecs.open(tmp_name, 'w+', out_encoding)

    with in_fh, out_fh:
        for line in fix_file(in_fh):
            out_fh.write(line)
    os.rename(tmp_name, in_name)


def purge_dir(directory):
    """This needs to raise exceptions."""
    for dirpath, dirnames, filenames in os.walk(directory):
        # We don't want to delete hidden filenames e.g. '.gitignore'
        filenames = [
            filename for filename in filenames if not filename[0] == '.'
        ]
        dirnames[:] = [
            dirname for dirname in dirnames if not dirname[0] == '.'
        ]
        for filename in filenames:
            os.unlink(os.path.join(dirpath, filename))
        for dirname in dirnames:
            shutil.rmtree(os.path.join(dirpath, dirname))


def ensure_exists(directory):
    """Ensure that a directory exists, creating it if it does not.

    Args:
        directory: A string of the desired directory path.

    Raises:
        OSError if any exception is raised on directory creation
        (other than an already exists type exception)."""
    try:
        os.makedirs(directory)
    except OSError as exception:
        if exception.errno != errno.EEXIST:
            raise


def all_zipped_files_startwith(zip_file, prefix):
    filenames = zip_file.namelist()
    for filename in filenames:
        extracted_path = os.path.abspath(os.path.join(prefix, filename))
        if not extracted_path.startswith(prefix):
            return False
    return True


def safely_unzip(destination, zipfile_path):
    """Safely unzip an archive into the destination path."""
    sucess = False
    with zipfile.ZipFile(zipfile_path, allowZip64=True) as zip_file:
        if all_zipped_files_startwith(zip_file, destination):
            zip_file.extractall(destination)
            success = True
        else:
            pass
    return success


def split_file_by_size(file_path, part_size_in_bytes):
    prefix = os.path.join(
        os.path.dirname(file_path),
        '%s.part.' % os.path.basename(file_path)
    )

    cmd = ['split', '-b%s' % part_size_in_bytes, file_path, prefix]
    subprocess.check_call(cmd)

    return sorted(glob.glob('%s*' % prefix))
