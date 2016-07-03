import argparse, shutil, sys, subprocess, os

packages = ["com/flickersearch", "com/flickersearch/authentication", "com/flickersearch/flicker", "com/flickersearch/history"]
test_packages = ["com/flickersearch/authentication", "com/flickersearch/flicker"]
cross_compilations = [
    ('linux', 'amd64'),
    ('windows', '386'),
]

files_to_remove = [
    'bin',
    'pkg'
]
def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('--clean', action='store_true', dest='clean',
                        help='Clean the project')
    parser.add_argument('--test', action='store_true', dest='test',
                       help='Runs tests')

    args = parser.parse_args()
    return args

def build_packages():
    for package in packages:
        gofmt(package)
        golint(package)
        govet(package)
        cmd = ['go', 'install', package]
        cmd_str = ' '.join(cmd)
        env_vars = os.environ.copy()
        env_vars['GOPATH'] = os.getcwd()
        for cross_compilation in cross_compilations:
            env_vars['GOOS'] = cross_compilation[0]
            env_vars['GOARCH'] = cross_compilation[1]
            if subprocess.call(cmd, env=env_vars) != 0:
                error_and_exit('Got a non-zero exit code while executing ' + cmd_str)

				
def run_tests(args):
    for package in test_packages:
		cmd = ['go', 'test', package, '-v']
		cmd_str = ' '.join(cmd)
		env_vars = os.environ.copy()
		env_vars['GOPATH'] = os.getcwd()
		if subprocess.call(cmd, env=env_vars) != 0:
			error_and_exit('Got a non-zero exit code while executing ' + cmd_str)

				
def clean():
    for f in files_to_remove:
        if os.path.exists(f):
            if os.path.isdir(f):
                shutil.rmtree(f)
            else:
                os.remove(f)	

def gofmt(pkg):
    cmd = ['go', 'fmt', pkg]
    cmd_str = ' '.join(cmd)
    if subprocess.call(cmd) != 0:
        error_and_exit('Got a non-zero exit code while executing ' + cmd_str)

def govet(pkg):
    cmd = ['go', 'vet', pkg]
    cmd_str = ' '.join(cmd)
    if subprocess.call(cmd) != 0:
        error_and_exit('Got a non-zero exit code while executing ' + cmd_str)

def golint(pkg):
    cmd = ['golint', pkg]
    cmd_str = ' '.join(cmd)
    if subprocess.call(cmd) != 0:
        error_and_exit('Got a non-zero exit code while executing ' + cmd_str)
		
def error_and_exit(msg):
    print 'Error:', msg
    sys.exit(1)
		
def main(args):
    if args.clean:
        clean()
    else:
        build_packages()
        if args.test:
            run_tests(args)
        
			
if __name__ == '__main__':
    args = parse_args()
    main(args)
   