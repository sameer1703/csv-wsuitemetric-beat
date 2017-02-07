from csvbeat import BaseTest

import os


class Test(BaseTest):

    def test_base(self):
        """
        Basic test with exiting Csvbeat normally
        """
        self.render_config_template(
                path=os.path.abspath(self.working_dir) + "/log/*"
        )

        csvbeat_proc = self.start_beat()
        self.wait_until( lambda: self.log_contains("csvbeat is running"))
        exit_code = csvbeat_proc.kill_and_wait()
        assert exit_code == 0
